import { AddForm } from "./components/AddForm";
import { TasksList } from "./components/TasksList";
import { Layout } from "./layouts/Layout";

function App() {
  return (
    <Layout>
      <section className="flex h-full flex-col gap-4">
        <section id="add-task">
          <AddForm />
        </section>
        <hr className="border-[0.5] border-neutral-400/50" />
        <section id="tasks-list">
          <TasksList />
        </section>
      </section>
    </Layout>
  );
}

export default App;
