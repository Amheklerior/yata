import { AddForm } from "./components/AddForm";
import { TasksList } from "./components/TasksList";
import { Layout } from "./layouts/Layout";

function App() {
  return (
    <Layout>
      <section className="grid grid-cols-1 gap-4">
        <section id="add-task">
          <AddForm />
        </section>
        <section id="tasks-list">
          <TasksList />
        </section>
      </section>
    </Layout>
  );
}

export default App;
