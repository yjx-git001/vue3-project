<template>
  <Transition name="toast">
    <div v-if="visible" class="toast-container">
      <div class="toast-content" :class="type">
        <svg v-if="type === 'success'" viewBox="0 0 24 24" class="toast-icon">
          <path d="M12 2C6.5 2 2 6.5 2 12S6.5 22 12 22 22 17.5 22 12 17.5 2 12 2M10 17L5 12L6.41 10.59L10 14.17L17.59 6.58L19 8L10 17Z" fill="currentColor"/>
        </svg>
        <svg v-else-if="type === 'error'" viewBox="0 0 24 24" class="toast-icon">
          <path d="M13 13H11V7H13M13 17H11V15H13M12 2A10 10 0 0 0 2 12A10 10 0 0 0 12 22A10 10 0 0 0 22 12A10 10 0 0 0 12 2Z" fill="currentColor"/>
        </svg>
        <svg v-else viewBox="0 0 24 24" class="toast-icon">
          <path d="M13,9H11V7H13M13,17H11V11H13M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z" fill="currentColor"/>
        </svg>
        <span class="toast-message">{{ message }}</span>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const visible = ref(false)
const message = ref('')
const type = ref<'success' | 'error' | 'info'>('info')

let timer: number | null = null

function show(msg: string, toastType: 'success' | 'error' | 'info' = 'info', duration = 2000) {
  message.value = msg
  type.value = toastType
  visible.value = true

  if (timer) {
    clearTimeout(timer)
  }

  timer = window.setTimeout(() => {
    visible.value = false
    timer = null
  }, duration)
}

defineExpose({
  show
})
</script>

<style scoped>
.toast-container {
  position: fixed;
  top: 20%;
  left: 50%;
  transform: translateX(-50%);
  z-index: 9999;
}

.toast-content {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px 24px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  min-width: 280px;
  max-width: 400px;
}

.toast-content.success {
  color: #10b981;
  border-left: 4px solid #10b981;
}

.toast-content.error {
  color: #ef4444;
  border-left: 4px solid #ef4444;
}

.toast-content.info {
  color: #3b82f6;
  border-left: 4px solid #3b82f6;
}

.toast-icon {
  width: 24px;
  height: 24px;
  flex-shrink: 0;
}

.toast-message {
  font-size: 15px;
  font-weight: 500;
  color: #333;
}

.toast-enter-active,
.toast-leave-active {
  transition: all 0.3s ease;
}

.toast-enter-from {
  opacity: 0;
  transform: translateX(-50%) translateY(-20px);
}

.toast-leave-to {
  opacity: 0;
  transform: translateX(-50%) translateY(-20px);
}
</style>
