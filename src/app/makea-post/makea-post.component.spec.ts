import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MakeaPostComponent } from './makea-post.component';

describe('MakeaPostComponent', () => {
  let component: MakeaPostComponent;
  let fixture: ComponentFixture<MakeaPostComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MakeaPostComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MakeaPostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
