#include "lib/Duck/DecoyDuck.h"
#include "lib/Duck/MallardDuck.h"
#include "lib/Duck/ModelDuck.h"
#include "lib/Duck/RedheadDuck.h"
#include "lib/Duck/RubberDuck.h"
#include "lib/DuckFunctions.h"
#include <cstdlib>

#include "lib/Duck/Fly/TriggeredQuackFlyBehavior.h"

int main()
{
	MallardDuck mallardDuck;
	auto originalFly = std::make_unique<FlyWithWings>();
	auto quack = std::make_unique<QuackBehavior>();

	auto triggeredQuackFly = std::make_unique<TriggeredQuackFlyBehavior>(
		std::move(originalFly),
		std::move(quack)
	);

	mallardDuck.SetFlyBehavior(std::move(triggeredQuackFly));
	PlayWithDuck(mallardDuck);
	PlayWithDuck(mallardDuck);
	PlayWithDuck(mallardDuck);
	PlayWithDuck(mallardDuck);


	RedheadDuck redheadDuck;
	PlayWithDuck(redheadDuck);
	PlayWithDuck(redheadDuck);
	PlayWithDuck(redheadDuck);
	PlayWithDuck(redheadDuck);

	RubberDuck rubberDuck;
	PlayWithDuck(rubberDuck);

	DecoyDuck decoyDuck;
	PlayWithDuck(decoyDuck);
	PlayWithDuck(decoyDuck);
	PlayWithDuck(decoyDuck);
	PlayWithDuck(decoyDuck);

	ModelDuck modelDuck;
	PlayWithDuck(modelDuck);

	modelDuck.SetFlyBehavior(std::make_unique<FlyWithWings>());
	PlayWithDuck(modelDuck);

	return EXIT_SUCCESS;
}