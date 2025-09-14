```mermaid
classDiagram
  class IDanceBehavior {
    <<Interface>>
    +void Dance()
  }
  
  class DanceMinuet {
    +void Dance()
  }
  
  class DanceWaltz {
    +void Dance()
  }
  
  class DanceNoWay {
    +void Dance()
  }
  
  class Duck {
    +void Quack()
    +void Swim()
    +void Fly()
    +void Dance()
    -m_flyBehavior: unique_ptr~IFlyBehavior~
    -m_quackBehavior: unique_ptr~IQuackBehavior~
    -m_danceBehavior: unique_ptr~IDanceBehavior~
  }
  
  class DecoyDuck {
    +void Display()
  }
  
  class MallardDuck {
    +void Display()
  }
  
  class ModelDuck {
    +void Display()
  }

  class RedheadDuck {
    +void Display()
  }
  
  class RubberDuck {
    +void Display()
  }

  DecoyDuck --|> Duck
  MallardDuck --|> Duck
  ModelDuck --|> Duck
  RedheadDuck --|> Duck
  RubberDuck --|> Duck

  DanceMinuet --|> IDanceBehavior
  DanceWaltz --|> IDanceBehavior
  DanceNoWay --|> IDanceBehavior

  IDanceBehavior --o Duck
