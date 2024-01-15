<template>
    <div class="form-landing">
        <div class="card">
            <div v-if="form">
                <header>
                    <h2>{{ form.name }}</h2>
                </header>
                <form @submit.prevent="submitPermit">
                    <template v-for="(q, i) in questions" :key="i">
                        <text-input
                            :label="q.label"
                            :ref="(el) => (questionRefs[q.id] = el)"
                            :type="q.type"
                            :autocomplete="q.autocomplete"
                            @keyup.enter="submitPermit"
                        />
                    </template>
                    <error-display :error="error"></error-display>
                    <c-button
                        class="w-100"
                    >
                        Submit
                    </c-button>
                </form>
            </div>
            <div v-else>
                Form could not be found. Please check if URL is correct.
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import TextInput from '@/components/global/TextInput.vue';
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

    constructor(
        id: string,
        label: string,
        ref: string,
        type = '',
        autocomplete = 'off'
    ) {
        this.id = id // should be same string as backend response / expected request
        this.label = label
        this.ref = ref
        this.type = type
        this.autocomplete = autocomplete
    }
}

const route = useRoute()

const form: Ref<IForm | null> = ref(null)
const permitResponse: Ref<IPermit | null> = ref(null)

// const questionRefs = ref({})
const questionRefs = ref<{[k: string]: InstanceType<typeof TextInput> }[]>([])

const questions = ref([
    new Question('first_name', 'First Name', 'firstName', 'given-name'),
    new Question('last_name', 'Last Name', 'lastName', 'family-name'),
    new Question('email', 'Email', 'email', 'email', 'email'),
    new Question('primary_phone', 'Phone', 'phone', 'tel', 'tel'),
    new Question('v_plate', 'Vehicle Plate', 'vPlate'),
    new Question('v_make', 'Vehicle Make', 'vMake'),
    new Question('v_Model', 'Vehicle Model', 'vModel'),
    new Question('v_color', 'Vehicle Color', 'vColor'),
    // new Question('duration', 'Duration', 'duration'),
])

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
            v_color: ''
        }
        questions.value.forEach((q) => {
            const val = questionRefs.value[q.id].value
            if (!val) {
                throw new Error(`${q.label} cannot be blank`)
            }
            payload[q.id] = val
        })
        permitResponse.value = await createPermit(form.value?.form_id, payload)
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
        // TODO submit permit
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
})
</script>
<style lang="scss" scoped>
.form-landing {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 1rem 0;
    background-color: var(--off-white-color);
    .card {
        max-width: 400px;
        width: 100%;
    }
}
</style>