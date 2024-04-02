<template>
  <div class="category-item-container">
    <span class="arrow-icon">
      <el-icon><ArrowRight /></el-icon>
    </span>
    <span class="doc-type-icon">
      <el-icon v-if="doc.docType == '1'"><Folder /></el-icon>
      <el-icon v-else-if="doc.docType == '2'"><Document /></el-icon>
    </span>
    <span class="doc-title">
      {{ doc.title }}
    </span>
  </div>
</template>

<script setup lang="ts">
import { PropType } from "vue";
import { models } from "../../../../wailsjs/go/models";
import { ArrowRight, Document, Folder } from "@element-plus/icons-vue";
const props = defineProps({
  doc: {
    type: Object as PropType<models.Doc>,
    required: true,
  },
});
</script>

<style lang="less" scoped>
.category-item-container {
  display: flex;
  align-items: center;
  height: 34px;
  .arrow-icon {
    width: 16px;
    height: 20px;
    margin-right: 5px;
  }
  .doc-type-icon {
    margin-right: 5px;
    color: var(--primary-color);
  }
  .doc-title {
    flex: 1;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    line-height: 20px;
    font-size: 14px;
    margin-top: -3px;
    color: var(--font-color-level1);
  }
}
</style>
