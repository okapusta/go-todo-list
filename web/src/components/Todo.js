import React from "react";

export const Todo = ({ id, createdAt, content }) => {
  return (
    <li key={ id }>{ content }</li>
  );
}
