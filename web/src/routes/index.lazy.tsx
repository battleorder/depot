import { Shell } from '@/components/shell'
import { createLazyFileRoute } from '@tanstack/react-router'

export const Route = createLazyFileRoute('/')({
  component: () => (
    <Shell>
      <h1 className="text-3xl font-bold">Welcome to BattleOrder.</h1>
      Hello, world!
    </Shell>
  ),
})
