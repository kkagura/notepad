<template>
  <div class="doc-page">
    <div :class="{ 'is-collapse': isLeftSideCollapsed }" class="left-side">
      <div class="left-header">
        <div class="left-title">文档中心</div>
        <el-icon
          class="clickable"
          @click="setLeftSideCollapsed(false)"
          v-if="isLeftSideCollapsed"
          ><Expand
        /></el-icon>
        <el-icon class="clickable" @click="setLeftSideCollapsed(true)" v-else
          ><Fold
        /></el-icon>
      </div>
      <CategoryList></CategoryList>
    </div>
    <div class="main-content"></div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import CategoryList from "./components/CategoryList.vue";
import { Fold, Expand } from "@element-plus/icons-vue";
const isLeftSideCollapsed = ref(false);

const setLeftSideCollapsed = (bool: boolean) => {
  isLeftSideCollapsed.value = bool;
};
</script>

<style lang="less" scoped>
.doc-page {
  display: flex;
  border: 1px solid var(--border-color-level1);
  .left-side {
    height: 100%;
    transition: width 0.3s ease;
    width: 280px;
    display: flex;
    flex-direction: column;
    border-right: 1px solid var(--border-color-level1);
    background-color: rgba(250, 251, 251, 1);
    color: var(--font-color-level1);
    &.is-collapse {
      width: 0;
      border-width: 0;
      overflow: hidden;
    }
    .left-header {
      display: flex;
      font-size: 14px;
      justify-content: space-between;
      align-items: center;
      height: 40px;
      padding: 0 10px;
      border-bottom: 1px solid var(--border-color-level1);
      overflow: hidden;
    }
  }
  .main-content {
    height: 100%;
    flex: 1;
    overflow: hidden;
  }
}
</style>
