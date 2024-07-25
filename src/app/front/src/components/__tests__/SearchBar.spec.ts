import { describe, it, expect } from 'vitest'

import { mount } from '@vue/test-utils'
import EmailComponent from '../EmailComponent.vue'

describe('EmailComponent', () => {
  it('renders properly', () => {
    const wrapper = mount(EmailComponent, { props: { msg: 'Hello Vitest' } })
    expect('Hello Vitest').toContain('Hello Vitest')
  })
})
