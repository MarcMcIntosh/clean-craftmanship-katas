
fn main() {
    println!("Hello, world!");
}



pub fn prime_factors(n: u64) -> Vec<u64> {
    let mut prime_numbers = Vec::<u64>::new();

    if n == 1 {
        prime_numbers.push(1);
        return prime_numbers;
    }

    let mut number = n; 
    let mut divisor = 2;

    while number > 1 {
        while number % divisor == 0 {
            prime_numbers.push(divisor);
            number /= divisor;
        }
        divisor += 1;
    }

    return prime_numbers;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    fn no_prime_factors_for_1() {
        let want = Vec::from([1]);
        let got = prime_factors(1);
        assert_eq!(want, got);
    }

    #[test]
    fn prime_factors_of_2_is_2() {
        let want = Vec::from([2]);
        let got = prime_factors(2);
        assert_eq!(want, got)
    }

    #[test]
    fn prime_factor_of_3_is_3() {
        let want = Vec::from([3]);
        let got = prime_factors(3);
        assert_eq!(want, got);
    }

    #[test]
    fn prime_factor_of_4_is_2_and_2() {
        let want = Vec::from([2, 2]);
        let got = prime_factors(4);
        assert_eq!(want, got);       
    }

    #[test] 
    fn prime_factor_of_5_is_5() {
        let want = Vec::from([5]);
        let got = prime_factors(5);
        assert_eq!(want, got);
    }

    #[test]
    fn prime_factor_of_6_is_2_and_3() {
        let want = Vec::from([2, 3]);
        let got = prime_factors(6);
        assert_eq!(want, got);
    }

    #[test]
    fn prime_factor_of_7_is_7() {
        let want = Vec::from([7]);
        let got = prime_factors(7);
        assert_eq!(want, got);
    }

    #[test]
    fn prime_factor_of_8_is_2_2_2() {
        // i really need to figure out how assert messages work 🤣
        let want = Vec::from([2, 2, 2]);
        let got = prime_factors(8);
        assert_eq!(want, got, "fuck!");
    }

    #[test]
    fn prime_factor_of_9_is_3_and_3() {
        let want = Vec::from([3, 3]);
        let got = prime_factors(9);
        assert_eq!(want, got)
    }

    #[test]
    fn prime_factors_of_primes() {
        let want = Vec::from([2,2,3,3,5,7,7,13]);
        let got = prime_factors(2 * 2 * 3 * 3 * 5 * 7 * 7 * 13);
        assert_eq!(want, got, "pretty sure it works now :)")
    }
}