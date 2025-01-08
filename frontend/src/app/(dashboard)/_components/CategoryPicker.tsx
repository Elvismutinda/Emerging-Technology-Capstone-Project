"use client";

import { TransactionType } from "@/types";
import React from "react";
import { useQuery } from "@tanstack/react-query";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Button } from "@/components/ui/button";
import {
  Command,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "@/components/ui/command";
import CreateCategoryDialog from "./CreateCategoryDialog";
import axios from "axios";

interface Props {
  type: TransactionType;
}

function CategoryPicker({ type }: Props) {
  const [open, setOpen] = React.useState(false);
  const [value, setValue] = React.useState("");
  const categoryQuery = useQuery({
    queryKey: ["categories", type],
    queryFn: async () => {
      const response = await axios.get(
        `http://localhost:8000/category/type/${type}`
      );

      return response.data;
    },
    staleTime: 1000 * 60 * 5,
  });

  const selectedCategory = categoryQuery.data?.find(
    (category: any) => category.name === value
  );
  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-[200px] justify-between"
        >
          {selectedCategory ? (
            <CategoryRow category={selectedCategory} />
          ) : (
            "Select category"
          )}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[200px] p-0">
        <Command
          onSubmit={(e) => {
            e.preventDefault();
          }}
        >
          <CommandInput
            placeholder="Search category..."
            onValueChange={(value) => setValue(value)}
          />
          <CommandGroup>
            {categoryQuery.data?.map((category: any) => (
              <CommandItem
                key={category.id}
                onSelect={() => {
                  setValue(category.name);
                  setOpen(false);
                }}
                className="w-full text-left"
              >
                <CategoryRow category={category} />
              </CommandItem>
            ))}
          </CommandGroup>
          <CreateCategoryDialog type={type} />
        </Command>
      </PopoverContent>
    </Popover>
  );
}

export default CategoryPicker;

function CategoryRow({ category }: { category: any }) {
  return (
    <div className="flex items-center gap-2">
      <span role="img">{category.icon}</span>
      <span>{category.name}</span>
    </div>
  );
}
