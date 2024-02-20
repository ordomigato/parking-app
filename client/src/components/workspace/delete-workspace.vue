<template>
    <header>
        <h2 class="danger-text">Danger Zone</h2>
    </header>
    <c-button
        :disabled="busy"
        @click="deleteModalOpen = true"
        danger
    >
        Delete Workspace
    </c-button>
    <ModalPopup
        v-if="deleteModalOpen"
    >
        <div class="modal-delete-container">
            <header>
                <h2>Delete workspace</h2>
            </header>
            <section>
                <p>Are you sure you wish to delete this workspace? This action cannot be undone.</p>
                <text-input
                    ref="confirmText"
                    label='Please type "confirm" to proceed'
                    @keyup.enter="onDeleteWorkspace"
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
                    @click="onDeleteWorkspace"
                    danger
                >
                    Delete Workspace
                </c-button>
            </footer>
        </div>
    </ModalPopup>
</template>
<script setup lang="ts">
import { routeNames } from '@/router/routeNames';
import { useWorkspaceStore } from '@/stores/workspaceStore';
import type { IWorkspace } from '@/types';
import { handleError } from '@/utils/error';
import { ref, type Ref } from 'vue';
import { useRouter } from 'vue-router';
import ModalPopup from '../common/modal-popup.vue';
import TextInput from '@/components/global/TextInput.vue'
import { useToastStore } from '@/stores/toastStore';
import { deleteWorkspace } from '@/services/workspace.service';

const props = defineProps({
    workspace: {
        type: Object as () => IWorkspace,
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

const onDeleteWorkspace = async () => {
    error.value = null
    busy.value = true
    try {
        if (confirmText.value?.value !== 'confirm') {
            throw new Error('Please type "confirm" to proceed')
        }
        await deleteWorkspace(props.workspace.workspace_id)
        workspaceStore.setActiveWorkspace(null)
        router.push({ name: routeNames.workspaces })
        toastStore.updateState("Successfully Deleted Workspace", "success")
    } catch (e) {
        error.value = handleError(e)
    } finally {
        busy.value = false
    }
}
</script>