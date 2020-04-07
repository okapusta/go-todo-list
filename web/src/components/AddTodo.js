import React from "react";

export const AddTodo = () => {
  return (
    <form>
      <span>
        <label>Add todo{ " " }
          <input name="todo"></input>
        </label>
      </span>
      <input type="submit" />
    </form>
  );
}
