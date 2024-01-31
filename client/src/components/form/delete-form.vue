<template>
    <header>
        <h2 class="danger-text">Danger Zone</h2>
    </header>
    <c-button
        :disabled="busy"
        @click="deleteModalOpen = true"
        danger
    >
        Delete Form
    </c-button>
    <ModalPopup
        v-if="deleteModalOpen"
    >
        <div class="modal-delete-container">
            <header>
                <h2>Delete Form</h2>
            </header>
            <section>
                <p>Are you sure you wish to delete this form? This action cannot be undone and all permits will be lost.</p>
                <text-input
                    ref="confirmText"
                    label='Please type "confirm" to proceed'
                    @keyup.enter="onDeleteForm"
                />
                <error-display :error="error"></error-display>
            </section>
            <footer>
                <c-button
                    :disabled="busy"
                    @click="deleteModalOpen = false; error = null"
                    danger
                    inverse
                >
                    Cancel
                </c-button>
                <c-button
                    :disabled="busy"
                    @click="onDeleteForm"
                    danger
                >
                    Delete Form
                </c-button>
            </footer>
        </div>
    </ModalPopup>
</template>
<script setup lang="ts">
import { routeNames } from '@/router/routeNames';
import { deleteForm } from '@/services/form.service';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import type { IForm } from '@/types';
import { handleError } from '@/utils/error';
import { ref, type Ref } from 'vue';
import { useRouter } from 'vue-router';
import ModalPopup from '../common/modal-popup.vue';
import TextInput from '@/components/global/TextInput.vue'
import { useToastStore } from '@/stores/toastStore';

const props = defineProps({
    form: {
        type: Object as () => IForm,
        required: true,
    }
})

const router = useRouter()
const workspaceStore = useWorkspaceStore()
const toastStore = useToastStore()

const deleteModalOpen = ref(false)
const confirmText = ref<InstanceType<typeof TextInput>>()

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const onDeleteForm = async () => {
    busy.value = true
    error.value = null
    try {
        if (!workspaceStore.currentWorkspace!) {
            throw new Error('something went wrong')
        }
        if (confirmText.value?.value !== 'confirm') {
            throw new Error('Please type "confirm" to proceed')
        }
        await deleteForm(workspaceStore.currentWorkspace?.workspace_id, props.form.form_id)
        router.push({ name: routeNames.forms })
        toastStore.updateState("Successfully Deleted Form", "success")
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}
</script>