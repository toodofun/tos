import { ref, onMounted, onUnmounted, watch } from 'vue'
import type { Ref } from 'vue'
import { type MaybeElementRef, useIntersectionObserver } from '@vueuse/core'

interface UseTimerOptions {
    root?: MaybeElementRef
    interval?: number
    immediate?: Boolean
}

function useTimer<T>(
    callback: () => Promise<T> | void,
    options: UseTimerOptions = {}
): {
    result: Ref<T | null>
    isFirstLoading: Ref<boolean>
} {
    const { interval = 1000, root = null, immediate = true } = options

    const isVisible = ref(false)
    useIntersectionObserver(root, ([{ isIntersecting }]) => {
        isVisible.value = isIntersecting
    })

    const result: Ref<T | null> = ref(null)
    const isFirstLoading: Ref<boolean> = ref(true)

    let timerId: number | null = null

    const call = async () => {
        result.value = await callback() || null
        isFirstLoading.value = false
        if (isVisible.value || root === null) {
            timerId = window.setTimeout(call, interval)
        }
    }

    const startTimer = async (): Promise<void> => {
        if (!timerId) {
            if (immediate) {
                await call()
            } else {
                timerId = window.setTimeout(call, interval)
            }
        }
    }

    const stopTimer = (): void => {
        if (timerId) {
            clearTimeout(timerId)
            timerId = null
        }
    }

    watch(isVisible, (newVal) => {
        if (newVal) {
            startTimer()
        } else {
            stopTimer()
        }
    })

    onMounted((): void => {
        if (isVisible.value || root === null) {
            startTimer()
        }
    })

    onUnmounted((): void => {
        stopTimer()
    })

    return { result, isFirstLoading }
}

export default useTimer
