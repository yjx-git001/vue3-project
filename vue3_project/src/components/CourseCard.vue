<template>
  <div :class="['course-card', `course-card--${layout}`]">
    <div class="course-card__image-wrapper">
      <img :src="course.image" :alt="course.title" class="course-card__image" />
      <span v-if="course.tag" class="course-card__tag">{{ course.tag }}</span>
    </div>
    <div class="course-card__info">
      <p class="course-card__title">{{ course.title }}</p>
      <div :class="['course-card__footer', `course-card__footer--${layout}`]">
        <div class="course-card__meta">
          <span class="course-card__original-price">¥{{ course.originalPrice }}</span>
          <span class="course-card__price">¥{{ course.price }}</span>
        </div>
        <button
          v-if="course.isPurchased"
          class="course-card__button course-card__button--learn"
          @click="handleClick"
        >
          立即学习
        </button>
        <button
          v-else
          class="course-card__button course-card__button--buy"
          @click="handleClick"
        >
          立即购买
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Course {
  id: number
  title: string
  tag?: string
  price: number
  originalPrice: number
  image: string
  isPurchased?: boolean
}

interface Props {
  course: Course
  layout?: 'horizontal' | 'grid'
}

const props = withDefaults(defineProps<Props>(), {
  layout: 'grid'
})

const emit = defineEmits<{
  click: [course: Course]
}>()

const handleClick = () => {
  emit('click', props.course)
}
</script>

<style scoped>
.course-card {
  background-color: #fff;
  border-radius: 10px;
  overflow: hidden;
}

.course-card--horizontal {
  flex: 0 0 200px;
}

.course-card--grid {
  padding: 8px;
}

.course-card__image-wrapper {
  position: relative;
}

.course-card__image {
  width: 100%;
  object-fit: cover;
}

.course-card--horizontal .course-card__image {
  height: 100px;
}

.course-card--grid .course-card__image {
  height: 120px;
  border-radius: 10px;
  margin-bottom: 8px;
}

.course-card__tag {
  position: absolute;
  top: 8px;
  left: 8px;
  font-size: 12px;
  color: #333;
  background-color: rgba(255,255,255,0.8);
  padding: 4px 8px;
  border-radius: 4px;
}

.course-card__info {
  padding: 8px;
}

.course-card--grid .course-card__info {
  padding: 0;
}

.course-card__title {
  font-size: 14px;
  font-weight: bold;
  margin: 0 0 8px 0;
}

.course-card__footer {
  display: flex;
}

.course-card__footer--horizontal {
  justify-content: space-between;
  align-items: flex-end;
}

.course-card__footer--grid {
  flex-direction: column;
}

.course-card__meta {
  display: flex;
  align-items: baseline;
  border-radius: 10px;
}

.course-card--horizontal .course-card__meta {
  flex-direction: column;
  align-items: flex-start;
}

.course-card--grid .course-card__meta {
  gap: 10px;
  margin-bottom: 10px;
}

.course-card__price {
  font-size: 20px;
  color: #ff6b35;
  font-weight: bold;
}

.course-card__original-price {
  font-size: 12px;
  color: #999;
  text-decoration: line-through;
}

.course-card--horizontal .course-card__original-price {
  margin-bottom: 4px;
}

.course-card__button {
  padding: 8px;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  font-size: 14px;
}

.course-card--horizontal .course-card__button {
  padding: 8px 20px;
  white-space: nowrap;
}

.course-card--grid .course-card__button {
  width: 85%;
  margin: 0 auto;
  display: block;
  border-radius: 15px;
}

.course-card__button--buy {
  background-color: #3b82f6;
  color: #fff;
}

.course-card__button--learn {
  background-color: #fff;
  color: #3b82f6;
  border: 1px solid #3b82f6;
}
</style>
