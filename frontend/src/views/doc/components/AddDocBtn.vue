<template>
  <el-dropdown>
    <div>
      <el-icon class="clickable">
        <Plus />
      </el-icon>
    </div>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item @click="handleAdd('2')">
          <el-icon><Document /></el-icon>
          文档
        </el-dropdown-item>
        <el-dropdown-item @click="handleAdd('1')">
          <el-icon><Folder /></el-icon>
          文件夹
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { Plus, Document, Folder } from "@element-plus/icons-vue";
const handleAdd = (docType: string) => {};
</script>

<style lang="postcss"></style>
