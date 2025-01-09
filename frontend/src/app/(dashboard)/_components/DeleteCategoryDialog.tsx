import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import useCurrentUser from "@/hooks/use-current-user";
import axios from "axios";
import { useRouter } from "next/navigation";
import React from "react";
import { toast } from "sonner";

interface Props {
  trigger: React.ReactNode;
  category: { id: string };
}

function DeleteCategoryDialog({ trigger, category }: Props) {
  const { token, user } = useCurrentUser();
  const userId = user?.id;
  const router = useRouter();

  const handleDeleteCategory = async () => {
    toast.loading("Deleting category...");

    try {
      const response = await axios.delete(
        `http://localhost:8000/category/${category.id}`,
        {
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
            "userId": userId,
          },
        }
      );

      console.log(response.data);

      toast.success("Category deleted successfully");
      router.refresh();
    } catch (error) {
      console.error("Error deleting category", error);
      toast.error("Error deleting category.");
    } finally {
      toast.dismiss();
    }
  };
  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>{trigger}</AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
          <AlertDialogDescription>
            This action cannot be undone. This will permanently delete your
            category.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction onClick={handleDeleteCategory}>
            Confirm
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}

export default DeleteCategoryDialog;
