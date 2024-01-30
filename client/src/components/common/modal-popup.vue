<template>
    <div class="modal-container">
        <div class="modal">
            <slot></slot>
        </div>
    </div>
</template>
<style setup lang="scss">
.modal-container {
    background-color: rgba(0,0,0,0.5);
    position: fixed;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: 9999;
    display: flex;
    align-items: center;
    justify-content: center;
    .modal {
        background-color: white;
        width: 400px;
        max-width: 100%;
        padding: 1.5rem;
        margin: 1rem;
        border-radius: 4px;
        box-shadow: 0 6px 12px rgba(0,0,0,0.1);
    }
}
</style>
