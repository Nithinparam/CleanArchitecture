package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Nithinparam/CleanArchitecture/controller"
	"github.com/Nithinparam/CleanArchitecture/repo"
	"github.com/Nithinparam/CleanArchitecture/router"
	"github.com/Nithinparam/CleanArchitecture/service"
)

var (
	repos          repo.PostRepo             = repo.NewFirestoreRepo()
	serve          service.PostService       = service.NewPostService(repos)
	postController controller.PostController = controller.NewPostController(serve)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./quickstart-1611068463766-firebase-adminsdk-114sr-3fc0c71eb8.json")
	const port string = ":8080"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Up...Running")
		fmt.Fprintln(w, "Hellooooo")
	})
	httpRouter.GET("/get", postController.GetPosts)
	httpRouter.GET("/add", postController.AddPosts)
	httpRouter.SERVE(port)

}
