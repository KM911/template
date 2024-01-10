<template>
  <a-space>
    <a-button type="primary">
      <template #icon>
        <icon-plus />
      </template>
    </a-button>
    <a-button type="primary">
      <template #icon>
        <icon-delete />
      </template>

      <template #default>Delete</template>
    </a-button>
  </a-space>
</template>

<script setup>
import { IconPlus, IconDelete } from '@arco-design/web-vue/es/icon';


</script>