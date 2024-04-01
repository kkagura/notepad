<template>
  <div class="app-container">
    <router-view v-slot="{ Component, route }">
      <template v-if="Component">
        <transition name="zoom-narrow">
          <keep-alive :max="10" :include="keepAlivePages">
            <suspense>
              <component :is="Component" :key="route.path"></component>
            </suspense>
          </keep-alive>
        </transition>
      </template>
    </router-view>
  </div>
</template>

<script setup lang="ts">
const keepAlivePages: string[] = [];
</script>

<style lang="less">
/* 缩放-放大效果 */
.zoom-narrow-enter-active,
.zoom-narrow-leave-active {
  transition: transform 0.35s, opacity 0.28s ease-in-out;
}
.zoom-narrow-enter-from {
  opacity: 0;
  transform: scale(0.97);
}
.zoom-narrow-leave-to {
  opacity: 0;
  transform: scale(1.03);
}
.app-container {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  & > div {
    height: 100%;
    overflow: auto;
  }
}
</style>
