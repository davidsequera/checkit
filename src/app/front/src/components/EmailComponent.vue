<template>
  <div class="flex flex-col px-2">
    <div class="mx-8">
      <h1 class="text-4xl font-bold mb-5">Inbox</h1>
      <SearchBar v-model="search" />
    </div>
    <div class="grid grid-cols-2 justify-items-center py-10">
      <EmailTable :search="search" @select-email="selectEmail" />
      <EmailBox :email="selectedEmail" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import EmailBox from './EmailBox.vue'
import SearchBar from './SearchBar.vue'
import EmailTable from './EmailTable.vue'
import type { Email } from '../models/Email'

const search = ref('')
const selectedEmail = ref<Email | null>(null)

const selectEmail = (email: Email) => {
  selectedEmail.value = email
}
</script>
