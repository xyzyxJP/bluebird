"use client";

import { ShelfTagSchema } from "@/app/shelf/(schema)/shelf-tag";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";
import { Form, FormControl, FormField, FormItem } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useToast } from "@/components/ui/use-toast";
import {
  CreateShelfTagDocument,
  CreateShelfTagMutation,
  CreateShelfTagMutationVariables,
  GetShelfTagsDocument,
  GetShelfTagsQuery,
} from "@/gql/gen/graphql";
import { useMutation, useQuery } from "@apollo/client";
import { zodResolver } from "@hookform/resolvers/zod";
import { ExclamationTriangleIcon } from "@radix-ui/react-icons";
import { useForm } from "react-hook-form";
import { TailSpin } from "react-loader-spinner";
import { z } from "zod";
import { ShelfTagColumns } from "../table/shelf-tag-columns";
import { ShelfTagTable } from "../table/shelf-tag-table";

const ShelfTagCreateFormSchema = z.object({
  name: z.string().min(1, { message: "Must be at least 1 character" }),
});

type ShelfTagCreateForm = z.infer<typeof ShelfTagCreateFormSchema>;

interface ShelfTagEditDialogProps {
  onOpenChange: (value: boolean) => void;
}

export function ShelfTagForm(props: ShelfTagEditDialogProps) {
  const { data, loading, error } =
    useQuery<GetShelfTagsQuery>(GetShelfTagsDocument);
  const form = useForm<ShelfTagCreateForm>({
    resolver: zodResolver(ShelfTagCreateFormSchema),
    mode: "onBlur",
  });
  const [
    createShelfTag,
    { loading: createShelfTagLoading, error: createShelfTagError },
  ] = useMutation<CreateShelfTagMutation, CreateShelfTagMutationVariables>(
    CreateShelfTagDocument,
    {
      refetchQueries: [{ query: GetShelfTagsDocument }],
    }
  );
  const { toast } = useToast();
  if (loading)
    return (
      <div className="h-screen flex items-center justify-center">
        <TailSpin />
      </div>
    );
  if (error)
    return (
      <Alert variant="destructive">
        <ExclamationTriangleIcon className="h-4 w-4" />
        <AlertTitle>Error</AlertTitle>
        <AlertDescription>{error.message}</AlertDescription>
      </Alert>
    );
  const tags = z.array(ShelfTagSchema).parse(data?.shelfTags ?? []);

  function onSubmit(data: ShelfTagCreateForm) {
    createShelfTag({
      variables: {
        name: data.name,
      },
    });
    toast({
      title: "タグを作成しました",
      description: data.name,
    });
  }

  return (
    <>
      <div className="mb-4">
        <ShelfTagTable columns={ShelfTagColumns} data={tags} />
      </div>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
          <div className="flex flex-1 items-end space-x-2">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                </FormItem>
              )}
            />
            <Button type="submit">追加</Button>
          </div>
        </form>
      </Form>
    </>
  );
}
