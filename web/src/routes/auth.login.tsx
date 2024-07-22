import { AuthShell } from '@/components/auth-shell'
import { Button } from '@/components/ui/button'
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { authStore } from '@/lib/auth'
import { supabase } from '@/lib/client'
import { zodResolver } from '@hookform/resolvers/zod'
import { AuthApiError, SignInWithPasswordCredentials } from '@supabase/supabase-js'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { createFileRoute, Link, redirect, useRouter } from '@tanstack/react-router'
import { useForm } from 'react-hook-form'
import { z } from 'zod'

const loginFormSchema = z.object({
  email: z.string().email('Invalid email'),
  password: z.string(),
});

export const Route = createFileRoute('/auth/login')({
  beforeLoad: () => {
    if (authStore.state.isAuthenticated) {
      throw redirect({
        to: '/',
      });
    }
  },
  component: () => {
    const router = useRouter();
    const client = useQueryClient();

    const form = useForm<z.infer<typeof loginFormSchema>>({
      resolver: zodResolver(loginFormSchema),
    });

    const { isPending, mutate: login } = useMutation({
      mutationFn: async (data: SignInWithPasswordCredentials) => {
        const res = await supabase.auth.signInWithPassword(data);
        if (res.error) {
          throw res.error;
        }
        return res;
      },
      onError: (err, _, __) => {
        if (err instanceof AuthApiError) {
          form.setError("email", {
            message: err.message,
            type: "validate",
          });
          return;
        }
      },
      onSuccess: () => {
        client.invalidateQueries({ queryKey: ['auth.user'] });
        router.navigate({ to: '/' });
      }
    });

    const onSubmit = (values: z.infer<typeof loginFormSchema>) => {
      login({ email: values.email, password: values.password });
    }

    return (
      <AuthShell
        title="Login"
        description="Enter your email below to login to your account"
      >
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="grid gap-4">
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input type="email" placeholder="john.doe@email.com" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )} />

            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Password</FormLabel>
                  <FormControl>
                    <Input type="password" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )} />

            <Button type="submit" disabled={isPending} className="w-full">
              Login
            </Button>
            <Button variant="outline" className="w-full">
              Login with Discord
            </Button>
          </form>
        </Form>
        <div className="mt-4 text-center text-sm">
          Don&apos;t have an account?{" "}
          <Link href="#" className="underline">
            Sign up
          </Link>
        </div>
      </AuthShell>
    );
  },
})
