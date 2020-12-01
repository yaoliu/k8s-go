//__author__ = "YaoYao"
//Date: 2020/9/27
package main

func CheckFunc(xx func() (string, error)) func() error {
	return func() error {
		obj, err := xx()
		if err != nil {
			return err
		}
		if obj != "ok" {
			return nil
		}
		return nil
	}
}

func main() {

}
