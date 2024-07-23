<template>
  <div class="flex justify-between">
    <EmailTable />
    <EmailBox />
  </div>
</template>

<script setup lang="ts">
import { defineComponent } from 'vue'
import EmailBox from './EmailBox.vue'
import EmailTable from './EmailTable.vue'
</script>

<script lang="ts">
export default defineComponent({
  name: 'EmailComponent'
})
</script>
