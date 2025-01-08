import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs";
import CategoryList from "./_components/CategoryList";
import { TabsContent } from "@radix-ui/react-tabs";
import AccountDetails from "./_components/AccountDetails";

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
      <Tabs defaultValue="account" className="container w-full pt-2">
        <TabsList className="grid w-full grid-cols-2">
          <TabsTrigger value="account">Account</TabsTrigger>
          <TabsTrigger value="categories">Categories</TabsTrigger>
        </TabsList>
        <TabsContent value="account">
          <div className="container flex p-4">
            {/* <AccountDetails /> */}
          </div>
        </TabsContent>
        <TabsContent value="categories">
          <div className="container flex flex-col gap-4 p-4">
            <CategoryList type="income" />
            <CategoryList type="expense" />
          </div>
        </TabsContent>
      </Tabs>
    </>
  );
}

export default page;
