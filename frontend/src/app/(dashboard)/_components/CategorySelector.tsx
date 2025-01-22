"use client";

import { TransactionType } from "@/types";
import React, { useEffect } from "react";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import useCurrentUser from "@/hooks/use-current-user";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import CreateCategoryDialog from "./CreateCategoryDialog";

interface Props {
  type: TransactionType;
  onChange: (value: string) => void;
}

function CategorySelector({ type, onChange }: Props) {
  const { token, user } = useCurrentUser();
  const userId = user?.id;

  const [value, setValue] = React.useState("");

  useEffect(() => {
    if (!value) return;
    onChange(value); // when the value changes, call the onChange callback
  }, [onChange, value]);

  const categoriesQuery = useQuery({
    queryKey: ["categories", type],
    queryFn: async () => {
      const response = await axios.get(
        `http://localhost:8000/category/type/${type}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
            userId: userId,
          },
        }
      );

      return response.data.data || [];
    },
  });

  const handleNewCategory = (newCategory: any) => {
    categoriesQuery.refetch();
    setValue(newCategory.name);
  };

  return (
    <div className="w-[200px]">
      <Select
        onValueChange={(selectedValue) => setValue(selectedValue)}
        value={value}
      >
        <SelectTrigger>
          <SelectValue placeholder="Select category" />
        </SelectTrigger>
        <SelectContent>
          <CreateCategoryDialog
            type={type}
            successCallback={handleNewCategory}
          />
          {categoriesQuery.data && categoriesQuery.data.length > 0 ? (
            categoriesQuery.data.map((category: any) => (
              <SelectItem key={category.name} value={category.name}>
                {category.name}
              </SelectItem>
            ))
          ) : (
            <p className="p-2 text-sm text-muted-foreground">
              No categories available
            </p>
          )}
        </SelectContent>
      </Select>
    </div>
  );
}

export default CategorySelector;

