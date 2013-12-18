package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"sync"
	"unsafe"
)

var dotproduct_code cu.Function

type dotproduct_args_t struct {
	arg_dst       unsafe.Pointer
	arg_prefactor float32
	arg_ax        unsafe.Pointer
	arg_ay        unsafe.Pointer
	arg_az        unsafe.Pointer
	arg_bx        unsafe.Pointer
	arg_by        unsafe.Pointer
	arg_bz        unsafe.Pointer
	arg_N         int
	argptr        [9]unsafe.Pointer
	sync.Mutex
}

var dotproduct_args dotproduct_args_t

func init() {
	dotproduct_args.argptr[0] = unsafe.Pointer(&dotproduct_args.arg_dst)
	dotproduct_args.argptr[1] = unsafe.Pointer(&dotproduct_args.arg_prefactor)
	dotproduct_args.argptr[2] = unsafe.Pointer(&dotproduct_args.arg_ax)
	dotproduct_args.argptr[3] = unsafe.Pointer(&dotproduct_args.arg_ay)
	dotproduct_args.argptr[4] = unsafe.Pointer(&dotproduct_args.arg_az)
	dotproduct_args.argptr[5] = unsafe.Pointer(&dotproduct_args.arg_bx)
	dotproduct_args.argptr[6] = unsafe.Pointer(&dotproduct_args.arg_by)
	dotproduct_args.argptr[7] = unsafe.Pointer(&dotproduct_args.arg_bz)
	dotproduct_args.argptr[8] = unsafe.Pointer(&dotproduct_args.arg_N)

}

// Wrapper for dotproduct CUDA kernel, asynchronous.
func k_dotproduct_async(dst unsafe.Pointer, prefactor float32, ax unsafe.Pointer, ay unsafe.Pointer, az unsafe.Pointer, bx unsafe.Pointer, by unsafe.Pointer, bz unsafe.Pointer, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	dotproduct_args.Lock()
	defer dotproduct_args.Unlock()

	if dotproduct_code == 0 {
		dotproduct_code = fatbinLoad(dotproduct_map, "dotproduct")
	}

	dotproduct_args.arg_dst = dst
	dotproduct_args.arg_prefactor = prefactor
	dotproduct_args.arg_ax = ax
	dotproduct_args.arg_ay = ay
	dotproduct_args.arg_az = az
	dotproduct_args.arg_bx = bx
	dotproduct_args.arg_by = by
	dotproduct_args.arg_bz = bz
	dotproduct_args.arg_N = N

	args := dotproduct_args.argptr[:]
	cu.LaunchKernel(dotproduct_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

var dotproduct_map = map[int]string{0: "",
	20: dotproduct_ptx_20,
	30: dotproduct_ptx_30,
	35: dotproduct_ptx_35}

const (
	dotproduct_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64


.visible .entry dotproduct(
	.param .u64 dotproduct_param_0,
	.param .f32 dotproduct_param_1,
	.param .u64 dotproduct_param_2,
	.param .u64 dotproduct_param_3,
	.param .u64 dotproduct_param_4,
	.param .u64 dotproduct_param_5,
	.param .u64 dotproduct_param_6,
	.param .u64 dotproduct_param_7,
	.param .u32 dotproduct_param_8
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<13>;
	.reg .s64 	%rd<23>;


	ld.param.u64 	%rd8, [dotproduct_param_0];
	ld.param.f32 	%f1, [dotproduct_param_1];
	ld.param.u64 	%rd9, [dotproduct_param_2];
	ld.param.u64 	%rd10, [dotproduct_param_3];
	ld.param.u64 	%rd11, [dotproduct_param_4];
	ld.param.u64 	%rd12, [dotproduct_param_5];
	ld.param.u64 	%rd13, [dotproduct_param_6];
	ld.param.u64 	%rd14, [dotproduct_param_7];
	ld.param.u32 	%r2, [dotproduct_param_8];
	cvta.to.global.u64 	%rd1, %rd8;
	cvta.to.global.u64 	%rd2, %rd14;
	cvta.to.global.u64 	%rd3, %rd13;
	cvta.to.global.u64 	%rd4, %rd12;
	cvta.to.global.u64 	%rd5, %rd11;
	cvta.to.global.u64 	%rd6, %rd10;
	cvta.to.global.u64 	%rd7, %rd9;
	.loc 1 11 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 12 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd7, %rd15;
	add.s64 	%rd17, %rd6, %rd15;
	add.s64 	%rd18, %rd5, %rd15;
	add.s64 	%rd19, %rd4, %rd15;
	add.s64 	%rd20, %rd3, %rd15;
	add.s64 	%rd21, %rd2, %rd15;
	.loc 1 14 1
	ld.global.f32 	%f2, [%rd19];
	.loc 1 13 1
	ld.global.f32 	%f3, [%rd16];
	.loc 1 14 1
	ld.global.f32 	%f4, [%rd20];
	.loc 1 13 1
	ld.global.f32 	%f5, [%rd17];
	.loc 1 15 1
	mul.f32 	%f6, %f5, %f4;
	fma.rn.f32 	%f7, %f3, %f2, %f6;
	.loc 1 14 1
	ld.global.f32 	%f8, [%rd21];
	.loc 1 13 1
	ld.global.f32 	%f9, [%rd18];
	.loc 1 15 1
	fma.rn.f32 	%f10, %f9, %f8, %f7;
	add.s64 	%rd22, %rd1, %rd15;
	.loc 1 15 1
	ld.global.f32 	%f11, [%rd22];
	fma.rn.f32 	%f12, %f10, %f1, %f11;
	st.global.f32 	[%rd22], %f12;

BB0_2:
	.loc 1 17 2
	ret;
}


`
	dotproduct_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64


.visible .entry dotproduct(
	.param .u64 dotproduct_param_0,
	.param .f32 dotproduct_param_1,
	.param .u64 dotproduct_param_2,
	.param .u64 dotproduct_param_3,
	.param .u64 dotproduct_param_4,
	.param .u64 dotproduct_param_5,
	.param .u64 dotproduct_param_6,
	.param .u64 dotproduct_param_7,
	.param .u32 dotproduct_param_8
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<13>;
	.reg .s64 	%rd<23>;


	ld.param.u64 	%rd8, [dotproduct_param_0];
	ld.param.f32 	%f1, [dotproduct_param_1];
	ld.param.u64 	%rd9, [dotproduct_param_2];
	ld.param.u64 	%rd10, [dotproduct_param_3];
	ld.param.u64 	%rd11, [dotproduct_param_4];
	ld.param.u64 	%rd12, [dotproduct_param_5];
	ld.param.u64 	%rd13, [dotproduct_param_6];
	ld.param.u64 	%rd14, [dotproduct_param_7];
	ld.param.u32 	%r2, [dotproduct_param_8];
	cvta.to.global.u64 	%rd1, %rd8;
	cvta.to.global.u64 	%rd2, %rd14;
	cvta.to.global.u64 	%rd3, %rd13;
	cvta.to.global.u64 	%rd4, %rd12;
	cvta.to.global.u64 	%rd5, %rd11;
	cvta.to.global.u64 	%rd6, %rd10;
	cvta.to.global.u64 	%rd7, %rd9;
	.loc 1 11 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 12 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd7, %rd15;
	add.s64 	%rd17, %rd6, %rd15;
	add.s64 	%rd18, %rd5, %rd15;
	add.s64 	%rd19, %rd4, %rd15;
	add.s64 	%rd20, %rd3, %rd15;
	add.s64 	%rd21, %rd2, %rd15;
	.loc 1 14 1
	ld.global.f32 	%f2, [%rd19];
	.loc 1 13 1
	ld.global.f32 	%f3, [%rd16];
	.loc 1 14 1
	ld.global.f32 	%f4, [%rd20];
	.loc 1 13 1
	ld.global.f32 	%f5, [%rd17];
	.loc 1 15 1
	mul.f32 	%f6, %f5, %f4;
	fma.rn.f32 	%f7, %f3, %f2, %f6;
	.loc 1 14 1
	ld.global.f32 	%f8, [%rd21];
	.loc 1 13 1
	ld.global.f32 	%f9, [%rd18];
	.loc 1 15 1
	fma.rn.f32 	%f10, %f9, %f8, %f7;
	add.s64 	%rd22, %rd1, %rd15;
	.loc 1 15 1
	ld.global.f32 	%f11, [%rd22];
	fma.rn.f32 	%f12, %f10, %f1, %f11;
	st.global.f32 	[%rd22], %f12;

BB0_2:
	.loc 1 17 2
	ret;
}


`
	dotproduct_ptx_35 = `
.version 3.2
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry dotproduct(
	.param .u64 dotproduct_param_0,
	.param .f32 dotproduct_param_1,
	.param .u64 dotproduct_param_2,
	.param .u64 dotproduct_param_3,
	.param .u64 dotproduct_param_4,
	.param .u64 dotproduct_param_5,
	.param .u64 dotproduct_param_6,
	.param .u64 dotproduct_param_7,
	.param .u32 dotproduct_param_8
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<13>;
	.reg .s64 	%rd<23>;


	ld.param.u64 	%rd8, [dotproduct_param_0];
	ld.param.f32 	%f1, [dotproduct_param_1];
	ld.param.u64 	%rd9, [dotproduct_param_2];
	ld.param.u64 	%rd10, [dotproduct_param_3];
	ld.param.u64 	%rd11, [dotproduct_param_4];
	ld.param.u64 	%rd12, [dotproduct_param_5];
	ld.param.u64 	%rd13, [dotproduct_param_6];
	ld.param.u64 	%rd14, [dotproduct_param_7];
	ld.param.u32 	%r2, [dotproduct_param_8];
	cvta.to.global.u64 	%rd1, %rd8;
	cvta.to.global.u64 	%rd2, %rd14;
	cvta.to.global.u64 	%rd3, %rd13;
	cvta.to.global.u64 	%rd4, %rd12;
	cvta.to.global.u64 	%rd5, %rd11;
	cvta.to.global.u64 	%rd6, %rd10;
	cvta.to.global.u64 	%rd7, %rd9;
	.loc 1 11 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 12 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB2_2;

	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd7, %rd15;
	add.s64 	%rd17, %rd6, %rd15;
	add.s64 	%rd18, %rd5, %rd15;
	add.s64 	%rd19, %rd4, %rd15;
	add.s64 	%rd20, %rd3, %rd15;
	add.s64 	%rd21, %rd2, %rd15;
	.loc 1 14 1
	ld.global.nc.f32 	%f2, [%rd19];
	.loc 1 13 1
	ld.global.nc.f32 	%f3, [%rd16];
	.loc 1 14 1
	ld.global.nc.f32 	%f4, [%rd20];
	.loc 1 13 1
	ld.global.nc.f32 	%f5, [%rd17];
	.loc 1 15 1
	mul.f32 	%f6, %f5, %f4;
	fma.rn.f32 	%f7, %f3, %f2, %f6;
	.loc 1 14 1
	ld.global.nc.f32 	%f8, [%rd21];
	.loc 1 13 1
	ld.global.nc.f32 	%f9, [%rd18];
	.loc 1 15 1
	fma.rn.f32 	%f10, %f9, %f8, %f7;
	add.s64 	%rd22, %rd1, %rd15;
	.loc 1 15 1
	ld.global.f32 	%f11, [%rd22];
	fma.rn.f32 	%f12, %f10, %f1, %f11;
	st.global.f32 	[%rd22], %f12;

BB2_2:
	.loc 1 17 2
	ret;
}


`
)