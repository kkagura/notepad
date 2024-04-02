<template>
  <div class="category-list-container">
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
    <div class="category-list">
      <CategoryItem v-for="doc in docTree" :doc="doc"></CategoryItem>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Search, Plus, Refresh } from "@element-plus/icons-vue";
import AddDocBtn from "./AddDocBtn.vue";
import { computed, ref } from "vue";
import { useRequest } from "@/hooks/useRequest";
import { GetDocTree } from "../../../../wailsjs/go/services/DocService";
import { models } from "../../../../wailsjs/go/models";
import CategoryItem from "./CategoryItem.vue";

const searchKey = ref("");
const requestContext = useRequest(GetDocTree);
const rawDocTree = ref<models.Doc[]>([]);
const docTree = computed(() => {
  return rawDocTree.value;
});
const fetchDocTree = () => {
  requestContext.request().then((res) => {
    rawDocTree.value = res;
  });
};
const onInit = () => {
  fetchDocTree();
};
onInit();
</script>

<style lang="less" scoped>
.category-list-container {
  height: 100%;
  width: 100%;
  flex: 1;
  overflow: hidden;
  padding: 0 10px;
  display: flex;
  flex-direction: column;
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
  .category-list {
    flex: 1;
    overflow-y: auto;
  }
}
</style>
