<template>
  <div class="category-list">
    <div class="search-wrap">
      <el-input
        v-model.trim="searchKey"
        :maxlength="20"
        clearable
        placeholder="搜索文档标题"
        :prefix-icon="Search"
      ></el-input>
    </div>
    <div class="category-tools">
      <span class="tools-title">目录</span>
      <el-icon class="clickable"><Refresh /></el-icon>
      <AddDocBtn></AddDocBtn>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Search, Plus, Refresh } from "@element-plus/icons-vue";
import AddDocBtn from "./AddDocBtn.vue";
import { ref } from "vue";
const searchKey = ref("");
</script>

<style lang="less" scoped>
.category-list {
  height: 100%;
  width: 100%;
  flex: 1;
  overflow-y: auto;
  padding: 0 10px;
  .search-wrap {
    padding: 10px 0;
  }
  .category-tools {
    margin-bottom: 10px;
    display: flex;
    color: var(--font-color-level2);
    .tools-title {
      flex: 1;
    }
    .el-icon {
      margin-left: 10px;
    }
  }
}
</style>
