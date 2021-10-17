import './App.css';

import {
  BrowserRouter as Router,
  Redirect,
  Switch,
  Route,
} from 'react-router-dom';

import Books from './pages/books';
import Book from './pages/book';
import AddBook from './pages/add-book';
import Layout from './components/layout';

function App() {
  return (
    <Router>
      <Layout>
        <Switch>
          <Route path="/" exact>
            <Books />
          </Route>
          <Redirect from="/books" to="/" exact />
          <Route path="/new" exact>
            <AddBook />
          </Route>
          <Route path="/books/:id" exact>
            <Book />
          </Route>
        </Switch>
      </Layout>
    </Router>
  );
}

export default App;
