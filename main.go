package main

import (
	"fmt"
	"net/http"

	"github.com/Nithinparam/CleanArchitecture/controller"
	"github.com/Nithinparam/CleanArchitecture/repo"
	"github.com/Nithinparam/CleanArchitecture/router"
	"github.com/Nithinparam/CleanArchitecture/service"
)

var (
	repos          repo.PostRepo             = repo.NewMysqlRepo()
	serve          service.PostService       = service.NewPostService(repos)
	postController controller.PostController = controller.NewPostController(serve)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {

	const port string = ":8080"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Up...Running")
		fmt.Fprintln(w, "Hellooooo")
	})
	httpRouter.GET("/get", postController.GetPosts)
	httpRouter.GET("/add", postController.AddPosts)
	httpRouter.SERVE(port)

}
