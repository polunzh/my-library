import './App.css';

import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

import Books from './pages/books';
import Book from './pages/book';
import AddBook from './pages/add-book';

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/" exact>
          <Books />
        </Route>
        <Route path="/add" exact>
          <AddBook />
        </Route>
        <Route path="/books/:id" exact>
          <Book />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
