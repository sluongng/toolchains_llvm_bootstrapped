#[cfg(all(target_arch = "aarch64", target_os = "linux"))]
extern "C" {
    fn __clear_cache(start: *mut u8, end: *mut u8);
}

#[cfg(all(target_arch = "aarch64", target_os = "linux"))]
#[inline(never)]
fn clear_cache(bytes: &mut [u8]) {
    unsafe {
        __clear_cache(bytes.as_mut_ptr(), bytes.as_mut_ptr().wrapping_add(bytes.len()));
    }
}

fn main() {
    #[cfg(all(target_arch = "aarch64", target_os = "linux"))]
    {
        let mut bytes = [0_u8; 1];
        clear_cache(&mut bytes);
    }

    println!("clear_cache linked");
}
