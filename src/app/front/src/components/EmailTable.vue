<template>
  <div class="max-w-4xl mx-auto w-11/12">
    <div class="overflow-auto border rounded-lg">
      <table>
        <thead>
          <tr class="hover:bg-muted/50">
            <th
              class="h-12 px-3 text-left align-middle font-medium text-muted-foreground cursor-pointer"
              @click="setSort('From')"
            >
              From
              <span v-if="sort.key === 'From'" class="ml-1">{{
                sort.order === 'asc' ? '↑' : '↓'
              }}</span>
            </th>
            <th
              class="h-12 px-3 text-left align-middle font-medium text-muted-foreground cursor-pointer"
              @click="setSort('Subject')"
            >
              Subject
              <span v-if="sort.key === 'Subject'" class="ml-1">{{
                sort.order === 'asc' ? '↑' : '↓'
              }}</span>
            </th>
            <th
              class="h-12 px-3 text-left align-middle font-medium text-muted-foreground cursor-pointer"
              @click="setSort('Date')"
            >
              Date
              <span v-if="sort.key === 'Date'" class="ml-1">{{
                sort.order === 'asc' ? '↑' : '↓'
              }}</span>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="email in emails"
            :key="email.MessageID"
            @click="selectEmail(email)"
            class="cursor-pointer hover:bg-gray-100"
          >
            <td class="p-3 align-middle font-medium">{{ email.From }}</td>
            <td class="p-3 align-middle">{{ email.Subject }}</td>
            <td class="p-3 align-middle">
              {{ new Date(email.Date).toLocaleDateString() }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import type { Email } from '../models/Email'

const props = defineProps<{ search: string }>()

const emit = defineEmits<{
  (e: 'selectEmail', email: Email): void
}>()

const selectEmail = (email: Email) => {
  emit('selectEmail', email)
}
const sort = ref({ key: 'Date', order: 'desc' })
const emails = ref<Email[]>([])

const fetchEmails = async (searchTerm: string = '') => {
  try {
    const response = await fetch(
      `http://localhost:8080/emails?search=${encodeURIComponent(searchTerm)}`
    )
    emails.value = await response.json()
    console.log(emails)
  } catch (error) {
    console.error('Error fetching emails:', error)
  }
}

onMounted(() => fetchEmails())

const setSort = (key: string) => {
  if (sort.value.key === key) {
    sort.value.order = sort.value.order === 'asc' ? 'desc' : 'asc'
  } else {
    sort.value = { key, order: 'asc' }
  }
  sortEmails()
}

const sortEmails = () => {
  emails.value.sort((a: any, b: any) => {
    const key = sort.value.key as keyof Email
    if (sort.value.order === 'asc') {
      return a[key] > b[key] ? 1 : -1
    } else {
      return a[key] < b[key] ? 1 : -1
    }
  })
}

watch(
  () => props.search,
  (newSearchTerm) => {
    fetchEmails(newSearchTerm)
  }
)

watch(sort, sortEmails)
</script>
