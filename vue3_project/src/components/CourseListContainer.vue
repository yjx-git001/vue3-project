<template>
  <div :class="['course-list-container', `course-list-container--${layout}`]">
    <slot />
  </div>
</template>

<script setup lang="ts">
interface Props {
  layout?: 'horizontal' | 'grid'
}

withDefaults(defineProps<Props>(), {
  layout: 'grid'
})
</script>

<style scoped>
.course-list-container--horizontal {
  display: flex;
  gap: 15px;
  overflow-x: auto;
}

.course-list-container--grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 15px;
}
</style>
