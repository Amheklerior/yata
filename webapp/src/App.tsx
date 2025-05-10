import { AddForm } from "./components/AddForm";
import { Separator } from "./components/Separator";
import { TasksList } from "./components/TasksList";
import { Layout } from "./layouts/Layout";

function App() {
  return (
    <Layout>
      <section className="flex h-full flex-col gap-4">
        <section id="add-task">
          <AddForm />
        </section>
        <Separator />
        <section id="tasks-list">
          <TasksList />
        </section>
      </section>
    </Layout>
  );
}

export default App;
