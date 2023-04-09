fn main() {
    println!("Hello, world!");
}

pub struct Stack {
    elements: Vec<i32>
}

impl Stack {
    fn new() -> Self {
        Self { elements: Vec::new() }
    }
    fn is_empty(&self) -> bool  {
        return self.get_size() == 0
    }

    fn get_size(&self) -> i32 {
        return self.elements.len().try_into().unwrap()
    }
    fn push(&mut self, n: i32) {
        // self.size += 1;
        // self.element = Some(n)
        self.elements.push(n)
    }
    fn pop(&mut self) -> Result<i32, &'static str> {
        if self.elements.len() == 0 {
            return Err("can not pop from empty stack")
        }
        match self.elements.pop()  {
            Some(n) => Ok(n),
            None => Err("can not pop from empty stack")
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4)
    }

    #[test]
    fn can_create_an_empty_stack() {
        let stack = Stack::new();
        assert!(stack.is_empty())
    }

    #[test]
    fn after_push_stack_is_not_empty() {
        let mut stack = Stack::new();
        stack.push(0);
        assert!(stack.is_empty() == false);
        assert!(stack.get_size() > 0);
    }

    #[test]
    fn after_one_push_and_one_pop_stack_is_empty() {
        let mut stack = Stack::new();
        stack.push(0);
        let _ = stack.pop();
        assert!(stack.is_empty())
    }

    #[test]
    fn after_two_pushes_size_should_be_two() {
        let mut stack = Stack::new();
        stack.push(0);
        stack.push(1);
        assert_eq!(stack.get_size(), 2)
    }

    #[test]
    fn popping_empty_stack_throws_error() {
       let mut stack = Stack::new();
       let result = stack.pop();
       let wanted = Err("can not pop from empty stack");
       assert_eq!(result, wanted)
    }

    #[test]
    fn after_pushing_x_it_will_pop_x() {
        let mut stack = Stack::new();
        stack.push(99);
        let result = stack.pop();
        let wanted = Ok(99);
        assert_eq!(result, wanted);
        stack.push(88);
        let result = stack.pop();
        let wanted = Ok(88);
        assert_eq!(result, wanted);
    }

    #[test]
    fn after_pushing_x_then_y_it_pop_should_return_y_then_x() {
        let mut stack = Stack::new();
        stack.push(88);
        stack.push(99);
        assert_eq!(stack.pop(), Ok(99));
        assert_eq!(stack.pop(), Ok(88));
    }
}