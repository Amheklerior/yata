import { Layout } from "./layouts/Layout";

function App() {
  return (
    <Layout>
      <section className="grid grid-cols-1 gap-4">
        <p> form for adding new tasks </p>
        <p> list of tasks </p>
      </section>
    </Layout>
  );
}

export default App;
