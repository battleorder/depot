import { Shell } from '@/components/shell'
import { createLazyFileRoute } from '@tanstack/react-router'

export const Route = createLazyFileRoute('/')({
  component: () => (
    <Shell>
      Hello, world!
    </Shell>
  ),
})
