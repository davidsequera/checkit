<template>
  <div class="w-full max-w-4xl mx-auto">
    <div class="flex items-center justify-between mb-4">
      <h1 class="text-2xl font-bold">Inbox</h1>
    </div>
    <SearchBar v-model="search" />
    <div class="overflow-auto border rounded-lg">
      <table>
        <thead>
          <tr>
            <th class="cursor-pointer" @click="setSort('From')">
              From
              <span v-if="sort.key === 'from'" class="ml-2">{{
                sort.order === 'asc' ? '↑' : '↓'
              }}</span>
            </th>
            <th class="cursor-pointer" @click="setSort('Subject')">
              Subject
              <span v-if="sort.key === 'subject'" class="ml-2">{{
                sort.order === 'asc' ? '↑' : '↓'
              }}</span>
            </th>
            <th class="cursor-pointer" @click="setSort('Date')">
              Date
              <span v-if="sort.key === 'date'" class="ml-2">{{
                sort.order === 'asc' ? '↑' : '↓'
              }}</span>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="email in filteredEmails" :key="email.id">
            <td class="font-medium">{{ email.From }}</td>
            <td>{{ email.Subject }}</td>
            <td>{{ new Date(email.Date).toLocaleDateString() }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed, onMounted, type Ref } from 'vue'
import SearchBar from './SearchBar.vue'
import { type Email } from '../models/Email'

export default defineComponent({
  name: 'InboxComponent',
  components: {
    SearchBar
  },
  setup() {
    const search = ref('')
    const sort = ref({ key: 'Date', order: 'desc' })
    const emails: Ref<Email[]> = ref([])

    // const emails = [
    //   { id: 1, from: 'olivia@example.com', subject: 'Meeting Agenda', date: '2023-06-01' },
    //   { id: 2, from: 'john@example.com', subject: 'Quarterly Report', date: '2023-05-15' },
    //   { id: 3, from: 'emily@example.com', subject: 'New Project Update', date: '2023-04-30' },
    //   { id: 4, from: 'michael@example.com', subject: 'Feedback Request', date: '2023-04-20' },
    //   { id: 5, from: 'sophia@example.com', subject: 'Invoice Reminder', date: '2023-03-28' },
    //   { id: 6, from: 'david@example.com', subject: 'Team Announcement', date: '2023-03-15' },
    //   { id: 7, from: 'isabella@example.com', subject: 'Upcoming Event', date: '2023-02-28' },
    //   { id: 8, from: 'alexander@example.com', subject: 'Budget Approval', date: '2023-02-10' },
    //   { id: 9, from: 'ava@example.com', subject: 'New Hire Announcement', date: '2023-01-25' },
    //   { id: 10, from: 'william@example.com', subject: 'Expense Report', date: '2023-01-05' }
    // ]
    const fetchEmails = async () => {
      try {
        const response = await fetch('http://localhost:8080/emails')
        const body = await response.json()
        emails.value = body
      } catch (error) {
        console.error('Error fetching emails:', error)
      }
    }

    onMounted(() => {
      fetchEmails()
    })

    const filteredEmails = computed(() => {
      return emails.value
        .filter(
          (email: Email) =>
            email.From.toLowerCase().includes(search.value.toLowerCase()) ||
            email.Subject.toLowerCase().includes(search.value.toLowerCase())
        )
        .sort((a: any, b: any) => {
          if (sort.value.order === 'asc') {
            return a[sort.value.key] > b[sort.value.key] ? 1 : -1
          } else {
            return a[sort.value.key] < b[sort.value.key] ? 1 : -1
          }
        })
    })

    const setSort = (key: string) => {
      if (sort.value.key === key) {
        sort.value.order = sort.value.order === 'asc' ? 'desc' : 'asc'
      } else {
        sort.value = { key, order: 'asc' }
      }
    }

    return {
      search,
      sort,
      filteredEmails,
      setSort
    }
  }
})
</script>
