/**
* Author : Samir Mehta
* Email: getsamir.mehta@gmail.com
**/

package devicedefaults

import (
	"android/soong/android"
	"android/soong/cc"
//        "fmt"   // required if we want to print usefull messages on the console
        "os"
)

func deviceFlags(ctx android.BaseContext) []string {
    var cflags []string
//    fmt.Fprintf(os.Stderr, "%s\n", "deviceFlags called") //example prints
//    fmt.Fprintf(os.Stderr, "%s\n", ctx.AConfig())
    product := os.Getenv("TARGET_PRODUCT") //currently using this for reference but you can refer to README for other variables.
//    fmt.Fprintf(os.Stderr, "%s\n", product)
    if product == "XYZ" {
       cflags = append(cflags, "-DIS_XYZ")
    }
    if product == "ABC" {
       cflags = append(cflags, "-DIS_ABC")
    }
    if product == "QWERTY" {
       cflags = append(cflags, "-DIS_QWERTY")
    }
    return cflags
}

func deviceDefaults(ctx android.LoadHookContext) {
    type props struct {
        Target struct {
            Android struct {
                Cflags  []string
                Enabled *bool
            }
            Host struct {
                Enabled *bool
            }
            Not_windows struct {
                Cflags []string
            }
        }
        Cflags []string
    }
    p := &props{}

    //p.Cflags = deviceFlags(ctx)
    p.Target.Android.Cflags = deviceFlags(ctx)
    ctx.AppendProperties(p)
}

/* This is called from the boot strap process
 * We register a module here
*/
func init() {
    android.RegisterModuleType("device_defaults", deviceDefaultFactory)
}

func deviceDefaultFactory() android.Module {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, deviceDefaults)
    return module
}
