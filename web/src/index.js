import React from "react";
import ReactDOM from "react-dom";

import { Container, Divider, Header } from 'semantic-ui-react';

import { TodoList } from './components/TodoList';

class App extends React.Component {
  render () {
    return (
      <Container>
        <Header>
          My todo list:
        </Header>
        <Divider />
        <TodoList />
      </Container>
    )
  }
}

ReactDOM.render(<App />, document.getElementById("root"));
