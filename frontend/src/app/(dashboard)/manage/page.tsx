import CategoryList from "./_components/CategoryList";

function page() {
  return (
    <>
      <div className="border-b bg-card">
        <div className="container flex flex-wrap items-center justify-between gap-6 py-8">
          <div>
            <p className="text-3xl font-bold">Manage</p>
            <p className="text-muted-foreground">
              Manage your category details here!
            </p>
          </div>
        </div>
      </div>
      <div className="container flex flex-col gap-4 p-4">
        <CategoryList type="INCOME" />
        <CategoryList type="EXPENSE" />
      </div>
    </>
  );
}

export default page;
