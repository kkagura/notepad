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
import {
  Add as AddDoc,
  GetDocTree,
} from "../../../../wailsjs/go/services/DocService";
import { models } from "../../../../wailsjs/go/models";
const handleAdd = async (docType: string) => {
  const res = await AddDoc(
    new models.Doc({
      id: "",
      docType,
      title: "未命名",
      parentId: "",
      children: [],
      createdAt: "",
      updatedAt: "",
      createdBy: "",
      updatedBy: "",
      fileId: "",
    })
  );
  console.log(res);
};
const onInit = async () => {
  GetDocTree().then((res) => {
    console.log(res);
  });
};
onInit();
</script>

<style lang="postcss" scoped></style>
