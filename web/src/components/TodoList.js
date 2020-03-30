import React from "react";

import { Todo } from "./Todo";

export class TodoList extends React.Component {
  constructor(props) {
    super(props);
    this.state = { todos: [], loading: false, error: null };
  }

  async fetch() {
    this.setState({ loading: true });
    try {
      const data = await fetch("http://localhost:8000/api/v1/todos")
        .then(resp => resp.json());// TODO: env, reactor me!!
      this.setState({ todos: data, loading: false })
    } catch (error) {
      this.setState({ error, loading: false })
    }
  }

  componentDidMount() {
    this.fetch();
  }

  render() {
    return (
      <ul>
        { this.state.todos && this.state.todos.map((todo) =>
            <Todo id={ todo.id } createdAt={ new Date() } content={ todo.content } />
          )
        }
      </ul>
    );
  }
}
