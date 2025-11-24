<template>
  <div class="tip-container">
    <span class="tip-prefix">Tips: 不知道输什么？不妨试试 </span>

    <!-- 只给数字加动画 -->
    <transition name="fade-slide" mode="out-in">
      <span :key="carriageDemo[index]" class="tip-number">
        {{ carriageDemo[index] }}
      </span>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const carriageDemo = ref(["92113", "081882", "T010192", "JC40124"])
const index = ref(0)

let switcher: number | undefined

onMounted(() => {
  switcher = window.setInterval(() => {
    index.value = (index.value + 1) % carriageDemo.value.length
  }, 2500)
})

onUnmounted(() => {
  clearInterval(switcher)
})
</script>

<style scoped>
.tip-container {
  display: inline-flex;
  align-items: center;
  padding: 6px 12px;
  font-size: 14px;
  line-height: 20px;
  border-radius: 6px;
  background: rgba(0, 0, 0, 0.04);
  color: #333;
  user-select: none;
  white-space: nowrap;
}

.tip-prefix {
  margin-right: 4px;
}

/* 只给数字元素固定宽度，防止数字长度变化导致跳动 */
.tip-number {
  display: inline-block;
  min-width: 60px;
  text-align: center;
  color: #4e9fff;
  font-weight: 600;
}


/* 动画 */
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.35s ease;
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(4px) scale(0.98);
}

.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(-4px) scale(0.98);
}
</style>
