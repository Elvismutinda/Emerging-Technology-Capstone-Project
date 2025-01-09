import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs";
import CategoryList from "./_components/CategoryList";
import { TabsContent } from "@radix-ui/react-tabs";

function page() {
  return (
    <>
      <div className="border-b bg-card">
        <div className="container flex flex-wrap items-center justify-between gap-6 py-8">
          <div>
            <p className="text-3xl font-bold">Manage</p>
            <p className="text-muted-foreground">
              Manage your account settings and categories
            </p>
          </div>
        </div>
      </div>
      <Tabs defaultValue="categories" className="container w-full pt-2">
        <TabsList className="grid w-full grid-cols-2">
          <TabsTrigger value="categories">Categories</TabsTrigger>
          <TabsTrigger value="account">Account</TabsTrigger>
        </TabsList>
        <TabsContent value="categories">
          <div className="container flex flex-col gap-4 p-4">
            <CategoryList type="INCOME" />
            <CategoryList type="EXPENSE" />
          </div>
        </TabsContent>
        <TabsContent value="account">
          <div className="container flex p-4">
            {/* <AccountDetails /> */}
          </div>
        </TabsContent>
      </Tabs>
    </>
  );
}

export default page;
