<template>
    <div class="form-landing">
        <div class="card">
            <div v-if="form">
                <header>
                    <h2>{{ form.name }}</h2>
                </header>
                <form
                    v-if="!permitResponse"
                    @submit.prevent="submitPermit"
                >
                    <template v-for="(q, i) in questions" :key="i">
                        <text-input
                            :label="q.label"
                            :ref="(el) => (questionRefs[q.id] = el)"
                            :type="q.type"
                            :min="q.min"
                            :max="q.max"
                            :autocomplete="q.autocomplete"
                            :defaultValue="q.defaultValue"
                        />
                    </template>
                    <label class="d-flex align-items-center justify-content-center py-4">
                        <input
                            type="checkbox"
                            v-model="rememberMe"
                        />
                        <span class="ms-2">Remember me?</span>
                    </label>
                    <error-display :error="error"></error-display>
                    <c-button
                        class="w-100"
                    >
                        Submit
                    </c-button>
                </form>
                <div v-else>
                    <div class="notice">
                        <p>Your parking permit has been created!</p>
                    </div>
                    <p v-if="permitResponse.expiry"><strong>Expiry: </strong>{{ convertDate(permitResponse.expiry) }} @ {{ convertTime(permitResponse.expiry) }}</p>
                    <p><strong>Plate Number: </strong>{{ permitResponse.v_plate }}</p>
                    <p><strong>Vehicle: </strong>{{ permitResponse.v_make }} {{ permitResponse.v_model }} ({{ permitResponse.v_color }})</p>
                    <p><strong>Name: </strong>{{ permitResponse.first_name }} {{ permitResponse.last_name }}</p>
                    <p><strong>Email: </strong>{{ permitResponse.email }}</p>
                    <p><strong>Phone: </strong>{{ permitResponse.primary_phone }}</p>
                    <p><strong>Created: </strong>{{ convertDate(permitResponse.created_at) }} @ {{ convertTime(permitResponse.created_at) }}</p>
                </div>
            </div>
            <div v-else>
                Form could not be found. Please check if URL is correct.
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import TextInput from '@/components/global/TextInput.vue';
import { convertDate, convertTime } from '@/utils/date';
import { getFormInfo } from '@/services/form.service';
import { createPermit } from '@/services/permit.service';
import type { IForm, IPermit, IPermitCreateRequest } from '@/types';
import { handleError } from '@/utils/error';
import { onMounted, ref, type Ref } from 'vue';
import { useRoute } from 'vue-router';

class Question {
    id: string
    label: string
    ref: string
    type: string
    autocomplete: string
    defaultValue: string
    min: number | undefined
    max: number | undefined

    constructor(
        id: string,
        label: string,
        ref: string,
        defaultValue = '',
        type = '',
        autocomplete = 'off'
    ) {
        this.id = id // should be same string as backend response / expected request
        this.label = label
        this.ref = ref
        this.type = type
        this.autocomplete = autocomplete
        this.defaultValue = defaultValue
    }

    setMin(num: number) {
        this.min = num
    }

    setMax(num: number) {
        this.max = num
    }

    setDefaultValue(val: string) {
        this.defaultValue = val
    }
}

const route = useRoute()

const form: Ref<IForm | null> = ref(null)
const permitResponse: Ref<IPermit | null> = ref(null)

// const questionRefs = ref({})
const questionRefs = ref<{[k: string]: InstanceType<typeof TextInput> }[]>([])

const questions: Ref<Question[]> = ref([])

const rememberMe: Ref<boolean> = ref(false)

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const submitPermit = async () => {
    busy.value = true
    error.value = null
    try {
        if (!form.value) {
            throw new Error("something went wrong")
        }
        // TODO submit permit!
        let payload: IPermitCreateRequest = {
            first_name: '',
            last_name: '',
            email: '',
            primary_phone: '',
            v_plate: '',
            v_make: '',
            v_model: '',
            v_color: '',
            duration: 1,
        }
        questions.value.forEach((q) => {
            const val = questionRefs.value[q.id].value
            if (!val) {
                throw new Error(`${q.label} cannot be blank`)
            }
            if (q.type === 'number') {
                payload[q.id] = parseInt(val)
            } else {
                payload[q.id] = val
            }
        })
        permitResponse.value = await createPermit(form.value?.form_id, payload)

        // save info locally for easier creation
        if (rememberMe.value) {
            localStorage.setItem(
                'customer_info', JSON.stringify(payload)
            )
        } else {
            localStorage.removeItem('customer_info')
        }
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}

const handleGetFormInfo = async () => {
    busy.value = true
    error.value = null
    try {
        const wpPath = route.params.workspacePath as string
        const formPath = route.params.formPath as string
        form.value = await getFormInfo(wpPath, formPath)
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }  
}

onMounted(async() => {
    await handleGetFormInfo()

    // handle remember me
    const cInfo = localStorage.getItem('customer_info')
    const c: IPermitCreateRequest = cInfo ? JSON.parse(cInfo) : null
    rememberMe.value = !!c

    questions.value = [
        new Question('first_name', 'First Name', 'firstName', c ? c.first_name : '', 'text', 'given-name'),
        new Question('last_name', 'Last Name', 'lastName', c ? c.last_name : '', 'text', 'family-name'),
        new Question('email', 'Email', 'email', c ? c.email : '', 'email', 'email'),
        new Question('primary_phone', 'Phone', 'phone', c ? c.primary_phone : '', 'tel', 'tel'),
        new Question('v_plate', 'Vehicle Plate', 'vPlate', c ? c.v_plate : '', 'text'),
        new Question('v_make', 'Vehicle Make', 'vMake', c ? c.v_make : '', 'text'),
        new Question('v_model', 'Vehicle Model', 'vModel', c ? c.v_model : '', 'text'),
        new Question('v_color', 'Vehicle Color', 'vColor', c ? c.v_color : '', 'text'),
    ]

    if (form.value?.cycle_data) {
        const durationQuestion = new Question('duration', `Duration (${form.value?.cycle_data.duration_limit.unit}) (Max: ${form.value.cycle_data.duration_limit.value})`, 'duration', '1', 'number')
        durationQuestion.setMin(1)
        durationQuestion.setMax(form.value.cycle_data.duration_limit.value)
        questions.value = [
            ...questions.value,
            durationQuestion
        ]
    }
})
</script>
<style lang="scss" scoped>
.form-landing {
    display: flex;
    min-height: 100vh;
    min-height: 100dvh;
    justify-content: center;
    padding: 1rem;
    background-color: var(--off-white-color);
    .card {
        max-width: 400px;
        width: 100%;
        height: 100%;
    }
}
</style>