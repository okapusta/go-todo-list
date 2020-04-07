import React from "react";
import ReactDOM from "react-dom";

import { Container, Divider, Header } from 'semantic-ui-react';

import { AddTodo } from './components/AddTodo';
import { TodoList } from './components/TodoList';

class App extends React.Component {
  render () {
    return (
      <Container>
        <Header>
          My todo list:
        </Header>
        <AddTodo />
        <Divider />
        <TodoList />
      </Container>
    )
  }
}

ReactDOM.render(<App />, document.getElementById("root"));
