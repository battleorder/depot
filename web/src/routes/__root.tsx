import React from 'react'
import { createRootRoute, Outlet } from '@tanstack/react-router'
import { updateAuth } from '@/lib/auth'

const RouterDevtools =
  process.env.NODE_ENV === 'production'
    ? () => null
    : React.lazy(() =>
      import('@tanstack/router-devtools').then((res) => ({
        default: res.TanStackRouterDevtools,
      })),
    )

export const Route = createRootRoute({
  beforeLoad: () => updateAuth(),
  component: () => (
    <>
      <Outlet />
      <RouterDevtools />
    </>
  ),
})
