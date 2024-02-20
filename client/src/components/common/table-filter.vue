<template>
    <div>
        <c-button
            @click="toggleModal"
        >
            Filter
        </c-button>
    </div>
    <ModalPopup
        v-if="openModal"
    >
        <h2>Visible Columns</h2>
        <div class="options">
            <label
                v-for="c in props.columns"
                :key="c.value"
            >
                <input
                    type="checkbox"
                    v-model="c.visible"
                />
                <span>{{ c.name }}</span>
            </label>
        </div>
        <footer>
            <c-button
                @click="toggleModal"
                inverse
            >
                Close
            </c-button>
            <c-button
                @click="handleSave"
            >
                Save
            </c-button>
        </footer>
    </ModalPopup>
</template>
<script setup lang="ts">
import { ref, type PropType } from 'vue';
import ModalPopup from './modal-popup.vue';
import type { IColumn } from '@/types';
import { saveLocal } from '@/utils/storage';

const props = defineProps({
    settingsId: {
        type: String,
        default: '',
    },
    columns: {
        type: Array as PropType<IColumn[]>,
        required: true,
    }
})

const openModal = ref(false)

const toggleModal = () => {
    openModal.value = !openModal.value
}

const handleSave = () => {
    // TODO set to localstorage
    if (!props.settingsId) {
        return
    }
    saveLocal(props.settingsId, JSON.stringify(props.columns))
    toggleModal()
}

</script>
<style lang="scss" scoped>
.options {
    margin: 1rem 0;
    label {
        display: block;
        input {
            margin-right: 0.5rem;
        }
    }
}
footer {
    display: flex;
    & > button {
        margin-right: 0.5rem;
    }
}
</style>