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
import { getUserId } from "@/hooks/get-user-id";

interface Props {
  type: TransactionType;
  trigger?: ReactNode;
}

function CreateCategoryDialog({ type, trigger }: Props) {
  const userId = getUserId();

  const [open, setOpen] = useState(false);
  const form = useForm<CreateCategoryRequest>({
    resolver: zodResolver(CreateCategorySchema),
    defaultValues: {
      type,
      name: "",
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
            Authorization: `Bearer ${localStorage.getItem("authToken")}`,
            "X-User-Id": userId,
          },
        }
      );

      console.log("Response ", response);

      toast.success("Category created successfully");

      setOpen(false);
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
            className="flex border-separate items-center justify-start rounded-none border-b px-3 py-3 text-muted-foreground"
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
                type === "income" ? "text-emerald-500" : "text-rose-500"
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
                    <Input
                      {...field}
                      placeholder="Enter category name"
                      defaultValue={""}
                    />
                  </FormControl>
                  {/* <FormDescription>
                    Transaction description (optional)
                  </FormDescription> */}
                </FormItem>
              )}
            />
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
              <Button type="submit" disabled={isLoading}>
                Save
                {isLoading && <Loader2 className="ml-2 h-4 w-4 animate-spin" />}
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}

export default CreateCategoryDialog;
