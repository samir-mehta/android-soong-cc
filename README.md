# android-soong-cc
Conditional Compilation based on some platform flags in Android for Soong/Ninja

As everyone is aware that android is movind towards soong from the original make based build system. 

<b>Soong</b>

Soong is the replacement for the old Android make-based build system. It replaces Android.mk files with Android.bp files, which are JSON-like simple declarative descriptions of modules to build.

You could convert your existing Android.mk to Android.bp via the androidmk tool available in src

Convert Android.mk files

Soong includes a tool perform a first pass at converting Android.mk files to Android.bp files:

<b>androidmk Android.mk > Android.bp</b>

Reference:
https://android.googlesource.com/platform/build/soong/

Now How do I write conditionals?

Soong deliberately does not support conditionals in Android.bp files. Instead, complexity in build rules that would require conditionals are handled in Go, where high level language features can be used and implicit dependencies introduced by conditionals can be tracked. Most conditionals are converted to a map property, where one of the values in the map will be selected and appended to the top level properties.

So the changes in this project partially support conditionals based on product variables.

Other list of device specific variables supported are 		"TARGET_BUILD_VARIANT", "TARGET_BUILD_APPS".





