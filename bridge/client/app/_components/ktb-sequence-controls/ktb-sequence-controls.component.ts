import { ChangeDetectionStrategy, Component, Input, ViewEncapsulation } from '@angular/core';
import { DataService } from '../../_services/data.service';
import { Sequence } from '../../_models/sequence';
import { SequenceStateControl } from '../../../../shared/models/sequence';
import { KtbConfirmationDialogComponent } from '../_dialogs/ktb-confirmation-dialog/ktb-confirmation-dialog.component';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'ktb-sequence-controls',
  templateUrl: './ktb-sequence-controls.component.html',
  styleUrls: ['./ktb-sequence-controls.component.scss'],
  host: {
    class: 'ktb-sequence-controls'
  },
  encapsulation: ViewEncapsulation.None,
  preserveWhitespaces: false,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class KtbSequenceControlsComponent {

  private _sequence?: Sequence;
  private _showLabel = false;
  public confirmationDialogRef?: MatDialogRef<KtbConfirmationDialogComponent>;

  @Input()
  get sequence(): Sequence | undefined {
    return this._sequence;
  }
  set sequence(sequence: Sequence | undefined) {
    if (this._sequence !== sequence) {
      this._sequence = sequence;
    }
  }

  @Input()
  get showLabel(): boolean {
    return this._showLabel;
  }
  set showLabel(value: boolean) {
    if (this._showLabel !== value) {
      this._showLabel = value;
    }
  }

  constructor(private dataService: DataService, public dialog: MatDialog) {
  }

  triggerResumeSequence(sequence: Sequence): void {
    this.dataService.sendSequenceControl(sequence, SequenceStateControl.RESUME);
  }

  triggerPauseSequence(sequence: Sequence): void {
    this.dataService.sendSequenceControl(sequence, SequenceStateControl.PAUSE);
  }

  triggerAbortSequence(sequence: Sequence): void {
    const data = {
      sequence,
      confirmCallback: (params: any) => {
        this.abortSequence(params.sequence);
      }
    };
    this.confirmationDialogRef = this.dialog.open(KtbConfirmationDialogComponent, {
      data,
    });
  }

  abortSequence(sequence: Sequence): void {
    this.dataService.sendSequenceControl(sequence, SequenceStateControl.ABORT);
  }
}
