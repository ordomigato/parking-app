<template>
    <header>
        <h2 class="danger-text">Danger Zone</h2>
    </header>
    <c-button
        :disabled="busy"
        @click="deleteModalOpen = true"
        danger
    >
        Delete Permit
    </c-button>
    <ModalPopup
        v-if="deleteModalOpen"
    >
        <div class="modal-delete-container">
            <header>
                <h2>Delete permit</h2>
            </header>
            <section>
                <p>Are you sure you wish to delete this permit? This action cannot be undone.</p>
                <text-input
                    ref="confirmText"
                    label='Please type "confirm" to proceed'
                    @keyup.enter="onDeletePermit"
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
                    :disabled="busy || confirmText?.value !== 'confirm'"
                    @click="onDeletePermit"
                    danger
                >
                    Delete Permit
                </c-button>
            </footer>
        </div>
    </ModalPopup>
</template>
<script setup lang="ts">
import { routeNames } from '@/router/routeNames';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import type { IFormattedPermit, IPermit } from '@/types';
import { handleError } from '@/utils/error';
import { ref, type Ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import ModalPopup from '../common/modal-popup.vue';
import TextInput from '@/components/global/TextInput.vue'
import { useToastStore } from '@/stores/toastStore';
import { deletePermit } from '@/services/permit.service';

const props = defineProps({
    permit: {
        type: Object as () => IPermit | IFormattedPermit,
        required: true,
    }
})

const route = useRoute()
const router = useRouter()
const workspaceStore = useWorkspaceStore()
const toastStore = useToastStore()

const deleteModalOpen = ref(false)
const confirmText = ref<InstanceType<typeof TextInput>>()

const error: Ref<Error | null> = ref(null)
const busy: Ref<boolean> = ref(false)

const onDeletePermit = async () => {
    error.value = null
    busy.value = true
    try {
        if (!workspaceStore.currentWorkspace!) {
            throw new Error('something went wrong')
        }
        if (confirmText.value?.value !== 'confirm') {
            throw new Error('Please type "confirm" to proceed')
        }

        await deletePermit(workspaceStore.currentWorkspace.workspace_id, route.params.id as string, props.permit.permit_id)
        router.push({ name: routeNames.form, params: { id: route.params.id } })
        toastStore.updateState("Successfully Deleted Permit", "success")
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}
</script>