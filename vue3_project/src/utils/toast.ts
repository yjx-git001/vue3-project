import { createApp } from 'vue'
import Toast from '@/components/Toast.vue'

let toastInstance: any = null

export function useToast() {
  if (!toastInstance) {
    const div = document.createElement('div')
    document.body.appendChild(div)
    const app = createApp(Toast)
    toastInstance = app.mount(div)
  }

  return {
    success: (message: string, duration?: number) => {
      toastInstance.show(message, 'success', duration)
    },
    error: (message: string, duration?: number) => {
      toastInstance.show(message, 'error', duration)
    },
    info: (message: string, duration?: number) => {
      toastInstance.show(message, 'info', duration)
    }
  }
}
