"use client";

import { TransactionType } from "@/types";
import React, { ReactNode, useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  CreateCategoryRequest,
  CreateCategorySchema,
} from "@/lib/validations/categories";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Loader2, PlusSquare } from "lucide-react";
import { cn } from "@/lib/utils";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import axios from "axios";
import { toast } from "sonner";
import useCurrentUser from "@/hooks/use-current-user";

interface Props {
  type: TransactionType;
  trigger?: ReactNode;
  successCallback: (category: any) => void;
}

function CreateCategoryDialog({ type, trigger, successCallback }: Props) {
  const { token, user } = useCurrentUser();
  const userId = user?.id;

  const [open, setOpen] = useState(false);
  const form = useForm<CreateCategoryRequest>({
    resolver: zodResolver(CreateCategorySchema),
    defaultValues: {
      category_type: type,
    },
  });

  const [isLoading, setIsLoading] = useState<boolean>(false);

  const handleCreateCategory = async (data: CreateCategoryRequest) => {
    setIsLoading(true);

    try {
      const response = await axios.post(
        "http://localhost:8000/category",
        data,
        {
          headers: {
            Authorization: `Bearer ${token}`,
            userId: userId,
          },
        }
      );

      toast.success(`Category ${data.name} created successfully`);

      successCallback(response.data.data);

      setOpen((prev) => !prev);
      form.reset();
    } catch (error) {
      console.error("Error creating category", error);
      toast.error("Failed to create the category.");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        {trigger ? (
          trigger
        ) : (
          <Button
            variant="ghost"
            className="flex w-full items-center justify-start rounded-none border-b px-3 py-3 text-muted-foreground"
          >
            <PlusSquare className="w-4 h-4 mr-2" />
            Create new
          </Button>
        )}
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>
            Create
            <span
              className={cn(
                "m-1",
                type === "INCOME" ? "text-emerald-500" : "text-rose-500"
              )}
            >
              {type}
            </span>
            category
          </DialogTitle>
        </DialogHeader>
        <Form {...form}>
          <form
            className="space-y-8"
            onSubmit={form.handleSubmit(handleCreateCategory)}
          >
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input placeholder="Enter category name" {...field} />
                  </FormControl>
                </FormItem>
              )}
            />
          </form>
        </Form>
        <DialogFooter>
          <DialogClose asChild>
            <Button
              type="button"
              variant="secondary"
              onClick={() => {
                form.reset();
              }}
            >
              Cancel
            </Button>
          </DialogClose>
          <Button
            onClick={form.handleSubmit(handleCreateCategory)}
            disabled={isLoading}
          >
            Save
            {isLoading && <Loader2 className="ml-2 h-4 w-4 animate-spin" />}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

export default CreateCategoryDialog;
